

import { ApiKey } from "@/types/api-key";
import { Response } from "@/schemas";

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