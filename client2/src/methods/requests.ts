import axios, { AxiosResponse } from "axios";
import {
	type T_FORMINFO_REGISTER,
	type T_APIRESULT_REGISTER,
	type T_FORMINFO_LOGIN,
	type T_APIRESULT_LOGIN,
	type T_TODO,
} from "../types";

// Routes
const ROUTE_PREFIX: string = import.meta.env.DEV ? "http://localhost:5000" : "";
const API_ROUTES = {
	LOGIN: ROUTE_PREFIX + "/api/login",
	REGISTER: ROUTE_PREFIX + "/api/register",
	VALIDATE: ROUTE_PREFIX + "/api/validateSession",
	GET_TODOS: ROUTE_PREFIX + "/api/getTodos",
};

export async function apiRequestLogin(
	formInfo: T_FORMINFO_LOGIN
): Promise<AxiosResponse<T_APIRESULT_LOGIN>> {
	try {
		return await axios<T_APIRESULT_LOGIN>({
			method: "POST",
			url: API_ROUTES.LOGIN,
			data: {
				...formInfo,
			},
		});
	} catch (err: any) {
		throw new Error(err);
	}
}

export async function apiRequestRegister(
	formInfo: T_FORMINFO_REGISTER
): Promise<AxiosResponse<T_APIRESULT_REGISTER>> {
	try {
		return await axios<T_APIRESULT_REGISTER>({
			method: "POST",
			url: API_ROUTES.REGISTER,
			data: {
				username: formInfo.username,
				password: formInfo.firstPassword,
			},
		});
	} catch (err: any) {
		throw new Error(err);
	}
}

export async function apiRequestGetTodos(
	username: string
): Promise<AxiosResponse<T_TODO[]>> {
	try {
		return await axios<T_TODO[]>({
			method: "POST",
			url: API_ROUTES.GET_TODOS,
			data: {
				username,
			},
		});
	} catch (err: any) {
		throw new Error(err);
	}
}

// export async function apiRequestValidateSession(
// 	userDataTokens: T_TOKENS
// ): Promise<AxiosResponse<T_APIRESULT_VALIDATE_ACCESS_TOKEN>> {
// 	console.log("Running apiRequestValidateSession");
// 	try {
// 		return await axios<T_APIRESULT_VALIDATE_ACCESS_TOKEN>({
// 			method: "POST",
// 			url: API_ROUTES.VALIDATE,
// 			headers: {
// 				Authorization: `Bearer ${userDataTokens.access_token}`,
// 			},
// 		});
// 	} catch (err: any) {
// 		throw new Error(err);
// 	}
// }

// interface I_APIREQUEST_MAKE_GUESS {
// 	allCurrentGuessInfo: T_ALL_CURRENT_GUESS_INFO;
// 	allCategoriesInfo: T_ALL_POSSIBLE_CATEGORIES_INFO;
// 	accessToken: string;
// 	gameSession: T_GAME_SESSION;
// }

// export async function apiRequestMakeGuess(
// 	paramObject: I_APIREQUEST_MAKE_GUESS
// ): Promise<AxiosResponse<string, any>> {
// 	try {
// 		const guessCategories = createRequestObjectFromCurrentGuessInfo(
// 			paramObject.allCurrentGuessInfo,
// 			paramObject.allCategoriesInfo
// 		);
// 		return await axios<string>({
// 			method: "POST",
// 			url: API_ROUTES.MAKE_GUESS,
// 			data: <T_APIREQUEST_MAKE_GUESS>{
// 				game_session_id: paramObject.gameSession.game_session_id,
// 				current_round: paramObject.gameSession.current_round,
// 				...guessCategories,
// 			},
// 			headers: {
// 				Authorization: `Bearer ${paramObject.accessToken}`,
// 			},
// 		});
// 	} catch (err: any) {
// 		throw new Error(err);
// 	}
// }

// export async function apiRequestGetPastGuesses(
// 	accessToken: string
// ): Promise<AxiosResponse<T_PAST_GUESSES, any>> {
// 	try {
// 		return await axios<T_PAST_GUESSES>({
// 			method: "POST",
// 			url: API_ROUTES.GET_PAST_GUESSES,
// 			headers: {
// 				Authorization: `Bearer ${accessToken}`,
// 			},
// 		});
// 	} catch (err: any) {
// 		throw new Error(err);
// 	}
// }
