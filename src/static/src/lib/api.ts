import { dev } from '$app/environment';

const API_HOST = !dev ? 'api:5004' : 'localhost:1323';

export const makeApiUrl = (pathname: string) => {
	return `http://${API_HOST}${pathname}`;
};
