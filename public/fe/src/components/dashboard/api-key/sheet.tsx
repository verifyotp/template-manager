'use client';
import { Button } from "@/registry/new-york/ui/button"
import { Input } from "@/registry/new-york/ui/input"
import { Label } from "@/components/ui/label"
import {
  Sheet,
  SheetContent,
  SheetDescription,
  SheetFooter,
  SheetHeader,
  SheetTitle,
  SheetTrigger,
} from "@/components/ui/sheet"
import * as React from "react"
import { useRouter } from 'next/navigation'
import { ApiKeyTable } from "@/components/dashboard/api-key/table"
import { useToast } from "@/components/ui/use-toast"
import { ApiKey } from "@/components/dashboard/api-key/table"


export function ApiKeySheet() {

  const [name, setName] = React.useState<string>("");
  const [data, setData] = React.useState<ApiKey[]>([]);
  const { toast } = useToast();
  const router = useRouter();

  async function handleFetchApiList() {
    try {
      const authToken = localStorage.getItem('authToken') as string;
      if (!authToken) {
        toast({
          title: "Error",
          description: "You are not authorized to view this page.",
        })
        router.push('/auth/login');
      }
      const response = await fetchApiList(authToken as string)
      setData(response.data as ApiKey[])
    } catch (error: any) {
      toast({
        variant: "destructive",
        title: "Error",
        description: error.message,
      })
    }
  }

  React.useEffect(() => {
   handleFetchApiList()
  }, []);

  const handleSubmit = () => {
    const authToken = localStorage.getItem('authToken') as string;
    if (!authToken) {
      router.push('/auth/login');
    }
    createApiKey(name, authToken)
      .then((data) => {
        toast({
          title: "Success",
          description: data.message,
        })
        handleFetchApiList()
      })
      .catch((error) => {
        toast({
          variant: "destructive",
          title: "Error",
          description: error.message,
        })
      });
  }


  const handleDelete = (id: string) => {
    const authToken = localStorage.getItem('authToken') as string;
    if (!authToken) {
      router.push('/auth/login');
    }
    deleteApiKey(id, authToken)
      .then((data) => {
        toast({
          title: "Success",
          description: data.message,
        })
        handleFetchApiList()
      })
      .catch((error) => {
        toast({
          variant: "destructive",
          title: "Error",
          description: error.message,
        })
      });
  }

  return (
    <Sheet>
      <SheetTrigger asChild>
        <Button variant="outline">Manage API Keys</Button>
      </SheetTrigger>
      <SheetContent className="w-[100%] sm:max-w-[80%] md:max-w-[50%]">
        <SheetHeader>
          <SheetTitle>API Keys Setting</SheetTitle>
          <SheetDescription>
            Manage your API keys.
          </SheetDescription>
        </SheetHeader>
        <div className="grid gap-4 py-4">
          <div className="grid grid-cols-4 items-center gap-4">
            <Label htmlFor="name" className="text-right">
              Name
            </Label>
            <Input
              id="name"
              value={name}
              onChange={(e) => setName(e.target.value)}
              placeholder="settlement-service"
              required className="col-span-3"
            />
          </div>
        </div>
        <SheetFooter>
          <div>
            <Button
              variant={"default"}
              onClick={handleSubmit}
              type="submit"
            >
              Create
            </Button>
          </div>
        </SheetFooter>
        <div className="py-6">
          <ApiKeyTable data={data} onDelete={handleDelete} />
        </div>
      </SheetContent>
    </Sheet>
  )
}

interface Response<T = any> {
  status: boolean;
  message: string;
  data?: T;
}

export async function createApiKey(name: string, authToken: string): Promise<Response> {
  const createApiKeyData = {
    name
  }
  const requestOptions = {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
    body: JSON.stringify(createApiKeyData)
  };

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/keys`, requestOptions);
    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }

    return data as Response;
  } catch (error: any) {
    throw new Error(error.message);
  }
}


export async function fetchApiList(authToken: string): Promise<Response<ApiKey[]>> {
  const requestOptions = {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
  };

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/keys`, requestOptions);
    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }

    return data as Response<ApiKey[]>;
  } catch (error: any) {
    throw new Error(error.message);
  }
}



export async function deleteApiKey(id: string, authToken: string): Promise<Response> {
  const requestOptions = {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${authToken}`
    },
  };

  try {
    const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/keys/${id}`, requestOptions);
    // Optionally handle response data here
    const data = await response.json();

    //check if the response is successful
    if (!data.status) {
      throw new Error(data.message);
    }

    return data as Response;
  } catch (error: any) {
    throw new Error(error.message);
  }
}