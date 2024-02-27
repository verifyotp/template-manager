'use server';

import {
    Response,
} from '@/schemas/auth';


export async function initiatePasswordReset(email: string,): Promise<Response> {
    const requestData = {
      email,
    };
  
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(requestData)
    };
  
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/users/reset-password`, requestOptions);
      // Optionally handle response data here
      const data = await response.json();
  
      //check if the response is successful
      if (!data.status) {
        throw new Error(data.message);
      }
      return data as Response;
    } catch (error : any) {
      throw new Error(error?.message);
    }
  }
  