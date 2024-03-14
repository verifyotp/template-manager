'use client';

import {
    Response,
} from '@/types';


export const logout = async ({token}: {token: string}) => {
    if (token) {
        const logoutData = {
            token
        };

        const requestOptions = {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(logoutData)
        };

        try {
            const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_BASE_URL}/api/users/logout`, requestOptions);
            const data = await response.json();
            localStorage.removeItem('authToken');
            return data as Response;
        } catch (error: any) {
            throw new Error(error?.message);
        }
    }
}