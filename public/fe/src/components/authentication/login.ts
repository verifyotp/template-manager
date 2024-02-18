"use server";

import { revalidatePath } from "next/cache";
import { z } from "zod";


export async function login(
  prevState: Response<LoginResponse> | { status: false },
  formData: FormData,
): Promise<Response<LoginResponse>> {

  console.log("login form data", formData.get("email"), formData.get("password"));


  const schema = z.object({
    email: z.string().min(1),
    password: z.string().min(1),
  });
  const parse = schema.safeParse({
    email: formData.get("email"),
    password: formData.get("password"),
  });



  if (!parse.success) {
    return { message: "Invalid form data", status: false };
  }

  const data = parse.data;

  try {
   const response = await loginUser(data.email, data.password);
    revalidatePath("/");
    return { message: `${response.message}`, status: response.status, data: response.data};
  } catch (e) {
    return { message: "Failed to login", status: false};
  }
}





export async function loginUser(email: string, password: string) : Promise<Response<LoginResponse>> {
    const loginData = {
      email,
      password,
    };
  
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(loginData)
    };
  
    try {
      const response = await fetch('http://localhost:9000/api/user/login', requestOptions);
      if (!response.ok) {
        const errorMessage = await response.text();
        throw new Error(errorMessage || 'Failed to login');
      }
  
      // Optionally handle response data here
      const data = await response.json();
  
      //check if the response is successful
      if (!data.status) {
        throw new Error(data.message);
      }
  
      // save token to local storage
      localStorage.setItem('session', data.session);
      localStorage.setItem('account', data.account);
      return data;
    } catch (error ) {
      throw new Error(error as string || 'An error occurred');
    } 
  }
  
  
  
  
  interface Session {
    id: string;
    account_id: string;
    device: any; // Define the Device type if needed
    token: string;
    expires_at: string;
    last_active: string;
    created_at: string;
  }
  
  interface Account {
    id: string;
    email: string;
    verified_at: string | null; // This can be a string or null
    created_at: string;
    updated_at: string | null; // This can be a string or null
  }
  
  interface LoginResponse {
    account: Account;
    session: Session;
  }
  
  interface Response<T = any> {
    status: boolean;
    message: string;
    data?: T;
  }
  
  