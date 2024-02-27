export interface Response<T = any> {
    status: boolean;
    message: string;
    data?: T;
}

