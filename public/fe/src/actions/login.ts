'use server';

import {
    LoginResponse 
} from '@/schemas/auth';

import {
    Response,
} from '@/schemas';

export async function login(email: string, password: string): Promise<Response<LoginResponse>> {
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
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/users/login`, requestOptions);
      // Optionally handle response data here
      const data = await response.json();
  
      //check if the response is successful
      if (!data.status) {
        throw new Error(data.message);
      }
      return data as Response<LoginResponse>;
    } catch (error: any) {
      throw new Error(error?.message);
    }
  }
  
  