'use client';
import { Button } from "@/components/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { FaRegCopy } from "react-icons/fa6";
import { IoMdEye, IoMdEyeOff } from "react-icons/io";
import { MdOutlineDelete } from "react-icons/md";
import * as React from "react"
import {
  ColumnDef,
  ColumnFiltersState,
  SortingState,
  VisibilityState,
  flexRender,
  getCoreRowModel,
  getFilteredRowModel,
  getPaginationRowModel,
  getSortedRowModel,
  useReactTable,
} from "@tanstack/react-table"
import { ArrowUpDown, ChevronDown } from "lucide-react";
import {
  DropdownMenu,
  DropdownMenuCheckboxItem,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { formatDate } from "@/lib/date";
import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"


interface AlertProps {
  Trigger: React.ReactNode;
  OnDelete: () => void;
}


export function AlertDialogDeleteApiKey({ Trigger, OnDelete }: AlertProps) {
  return (
    <AlertDialog>
      <AlertDialogTrigger asChild>
        {Trigger}
      </AlertDialogTrigger>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>Are you absolutely sure?</AlertDialogTitle>
          <AlertDialogDescription>
            This action cannot be undone. This will <b>permanently delete</b> this
            api key.
          </AlertDialogDescription>
        </AlertDialogHeader>
        <AlertDialogFooter>
          <AlertDialogCancel
          >Cancel</AlertDialogCancel>
          <AlertDialogAction
            onClick={() => {
              OnDelete()
            }}
          >Continue</AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}


interface getColumnProps {
  OnDelete: (id: string) => void;
}

function getColumns({ OnDelete }: getColumnProps): ColumnDef<ApiKey>[] {

  return [
    {
      accessorKey: "name",
      header: "Name",
      cell: ({ row }) => (
        <div >{row.getValue("name")}</div>
      ),
    },
    {
      accessorKey: "created_at",
      header: ({ column }) => {
        return (
          <Button
            variant="ghost"
            onClick={() => column.toggleSorting(column.getIsSorted() === "desc")}
          >
            Created At
            <ArrowUpDown className="ml-2 h-4 w-4" />
          </Button>
        )
      },
      cell: ({ row }) => {
        const createdAt = row.getValue("created_at")
        const formatDatetime = formatDate(createdAt as Date);
        return <div className="text-center">{formatDatetime}</div>
      },
    },
    {
      accessorKey: "secret",
      header: () => <div className=" password">Secret</div>,
      cell: function Cell ({ row }) {
        const secret = row.getValue("secret") as string

        const [hiddenSecret, setHiddenSecret] = React.useState<boolean>(false)
        const [clicked, setClicked] = React.useState<boolean>(false)

        const copySecret = () => navigator.clipboard.writeText(secret as string);
        const maskSecret = (secret: string) => (secret.length < 10) ? secret : secret.slice(0, 7) + "*****" + secret.slice(-5);

        const handleSecret = () => setHiddenSecret(!hiddenSecret)
        const handleCopy = () => { setClicked(true); copySecret(); setTimeout(() => { setClicked(false) }, 500) }


        return <div className="font-medium flex items-center justify-center space-x-2 gap-2">
          <div >
            {hiddenSecret ? maskSecret(secret) : "********************"}
          </div>
          <div>
            <button onClick={handleSecret}>
              {
                hiddenSecret ? < IoMdEyeOff /> : <IoMdEye className="opacity-[0.6] hover:opacity-[1]" />
              }
            </button>
          </div>
          <div className="">
            <button onClick={handleCopy}>
              <FaRegCopy
                className={`opacity-[0.6] hover:opacity-[1] ${clicked ? "text-[13px]-duration-[0.4s]" : "text-[15px]-duration-[0.4s]"}`}
              />
            </button>
          </div>
        </div>
      },
    },
    {
      accessorKey: "id",
      header: "",
      cell: ({ row }) => {

        const id = row.getValue("id") as string
        const handleDelete = () => {
          OnDelete(id)
        }

        return <div>
          <AlertDialogDeleteApiKey
            OnDelete={handleDelete}
            Trigger={
              <Button variant="outline" className="hover:bg-red-500 hover:text-black">
                <MdOutlineDelete />
              </Button>
            } />
        </div>
      },
    },
  ]
}

export type ApiKey = {
  id: string
  name: string
  secret: string
  created_at: string
}

interface ApiKeyTableProps {
  data: ApiKey[]
  onDelete: (id: string) => void;
}

export function ApiKeyTable({ data, onDelete }: ApiKeyTableProps) {
  const [sorting, setSorting] = React.useState<SortingState>([])
  const [columnFilters, setColumnFilters] = React.useState<ColumnFiltersState>([])
  const [columnVisibility, setColumnVisibility] = React.useState<VisibilityState>({});
  const [rowSelection, setRowSelection] = React.useState({});

  const columns = getColumns({ OnDelete: onDelete })

  const table = useReactTable({
    data: data,
    columns,
    onSortingChange: setSorting,
    onColumnFiltersChange: setColumnFilters,
    getCoreRowModel: getCoreRowModel(),
    getPaginationRowModel: getPaginationRowModel(),
    getSortedRowModel: getSortedRowModel(),
    getFilteredRowModel: getFilteredRowModel(),
    onColumnVisibilityChange: setColumnVisibility,
    onRowSelectionChange: setRowSelection,
    state: {
      sorting,
      columnFilters,
      columnVisibility,
      rowSelection,
    },
  })


  return (
    <div className="w-full">
      <div className="flex items-center py-4">
        <Input
          placeholder="Filter names..."
          value={(table.getColumn("name")?.getFilterValue() as string) ?? ""}
          onChange={(event) =>
            table.getColumn("name")?.setFilterValue(event.target.value)
          }
          className="max-w-sm"
        />
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" className="ml-auto">
              Columns <ChevronDown className="ml-2 h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            {table
              .getAllColumns()
              .filter((column) => column.getCanHide())
              .map((column) => {
                return (
                  <DropdownMenuCheckboxItem
                    key={column.id}
                    className="capitalize"
                    checked={column.getIsVisible()}
                    onCheckedChange={(value) =>
                      column.toggleVisibility(!!value)
                    }
                  >
                    {column.id}
                  </DropdownMenuCheckboxItem>
                )
              })}
          </DropdownMenuContent>
        </DropdownMenu>
      </div>
      <div className="rounded-md border">
        <Table >
          <TableHeader>
            {table.getHeaderGroups().map((headerGroup) => (
              <TableRow className="text-center" key={headerGroup.id}>
                {headerGroup.headers.map((header) => {
                  return (
                    <TableHead className="text-center" key={header.id}>
                      {header.isPlaceholder
                        ? null
                        : flexRender(
                          header.column.columnDef.header,
                          header.getContext()
                        )}
                    </TableHead>
                  )
                })}
              </TableRow>
            ))}
          </TableHeader>
          <TableBody>
            {table.getRowModel().rows?.length ? (
              table.getRowModel().rows.map((row) => (
                <TableRow
                  key={row.id}
                  data-state={row.getIsSelected() && "selected"}
                >
                  {row.getVisibleCells().map((cell) => (
                    <TableCell className="text-center" key={cell.id}>
                      {flexRender(
                        cell.column.columnDef.cell,
                        cell.getContext()
                      )}
                    </TableCell>
                  ))}
                </TableRow>
              ))
            ) : (
              <TableRow>
                <TableCell
                  colSpan={columns.length}
                  className="h-24 text-center"
                >
                  No results.
                </TableCell>
              </TableRow>
            )}
          </TableBody>
        </Table>
      </div>
      <div className="flex items-center justify-end space-x-2 py-4">
        <div className="space-x-2">
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.previousPage()}
            disabled={!table.getCanPreviousPage()}
          >
            Previous
          </Button>
          <Button
            variant="outline"
            size="sm"
            onClick={() => table.nextPage()}
            disabled={!table.getCanNextPage()}
          >
            Next
          </Button>
        </div>
      </div>
    </div>
  )
}


