import Axios, { AxiosError, type AxiosResponse, type InternalAxiosRequestConfig } from "axios";

type ResponseFnType = (response: AxiosResponse<any, any>) => void;
type RequestFnType = (response: InternalAxiosRequestConfig) => void;
type ErrorResponseFnType = (error: AxiosError) => Promise<void>;

const responses: ResponseFnType[] = [];
const requestes: RequestFnType[] = [];
const errorsResponses: ErrorResponseFnType[] = [];

const api = Axios.create({
	baseURL: API_URL,
});

export function OnResponse(action: ResponseFnType) {
	responses.push(action);
}

export function OnRequest(action: RequestFnType) {
	requestes.push(action);
}

export function OnError(action: ErrorResponseFnType) {
	errorsResponses.push(action);
}

api.interceptors.response.use(
	(response: AxiosResponse) => {
		responses.forEach((action) => action(response));
		return response;
	},
	async (error: AxiosError) => {
		for (const action of errorsResponses) {
			await action(error);
		}
		return Promise.reject(error);
	},
);

api.interceptors.request.use((config: InternalAxiosRequestConfig) => {
	requestes.forEach((action) => action(config));
	return config;
});

export { api };
