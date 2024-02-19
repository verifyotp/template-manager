'use client';

import {  format } from 'date-fns';
export function formatDate(date: Date): string {
    return format(date, "E yyyy-MM-dd HH a");
}