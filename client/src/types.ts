export type T_FORMINFO_REGISTER = {
	username: string;
	firstPassword: string;
	secondPassword: string;
};

export type T_APIRESULT_VALID = {
	valid: boolean;
};

export type T_FORMINFO_LOGIN = {
	username: string;
	password: string;
};

export type T_APIRESULT_LOGIN = {
	valid: boolean;
};

export type T_TODO = {
	id: number;
	title: string;
	content: string;
};
