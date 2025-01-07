import ax from "axios";

export const baseUrl = import.meta.env.VITE_BACKEND_PREFIX;

export const axios = ax.create({
    baseURL: baseUrl,
    withCredentials: true,
});
