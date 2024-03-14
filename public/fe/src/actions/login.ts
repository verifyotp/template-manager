'use server';

import {
    LoginResponse 
} from '@/types/auth';

import {
    Response,
} from '@/types';

export async function login({email, password}: {
  email: string;
  password: string;
}): Promise<Response<LoginResponse>> {

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
      return data as Response<LoginResponse>;
    } catch (error: any) {
      throw new Error(error?.message);
    }
  }
  
  