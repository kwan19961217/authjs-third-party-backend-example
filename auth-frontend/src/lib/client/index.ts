import axios, { type AxiosInstance } from 'axios';
import { AUTH_BACKEND_URL } from '$env/static/private';

// client for the auth adapter
export default function authClient(): AxiosInstance {
	return axios.create({
		baseURL: AUTH_BACKEND_URL,
		timeout: 1000
	});
}
