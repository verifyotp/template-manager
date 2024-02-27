export interface Session {
    id: string;
    account_id: string;
    device: any; // Define the Device type if needed
    token: string;
    expires_at: string;
    last_active: string;
    created_at: string;
}

export interface Account {
    id: string;
    email: string;
    verified_at: string | null; // This can be a string or null
    created_at: string;
    updated_at: string | null; // This can be a string or null
}

export interface LoginResponse {
    account: Account;
    session: Session;
}
