import styles from "./Login.module.scss";
import { useState } from "react";
import { type T_APIRESULT_VALID, type T_FORMINFO_LOGIN } from "../../types";
import { apiRequestLogin } from "../../methods/requests";
import { AxiosResponse } from "axios";

const Login = () => {
  const [userFormInfo, setUserFormInfo] = useState<T_FORMINFO_LOGIN>({
    username: "",
    password: "",
  });

  function inputHandler(e: React.ChangeEvent<HTMLInputElement>): void {
    setUserFormInfo((userFormInfo: T_FORMINFO_LOGIN) => {
      return {
        ...userFormInfo,
        [e.target.name]: e.target.value,
      };
    });
  }

  async function login(): Promise<void> {
    const res: AxiosResponse<T_APIRESULT_VALID> = await apiRequestLogin(
      userFormInfo
    );
    console.log(`login valid: ${res.data.valid}`);
  }

  return (
    <div className={`card mx-auto mt-3 ${styles.root}`}>
      <div className="card-body text-center">
        <h5 className="card-title text-center">
          <label htmlFor="username" className="mb-1">
            Username
          </label>
          <input
            className="form-control"
            onChange={inputHandler}
            value={userFormInfo.username}
            type="text"
            id="username"
            name="username"
          />
        </h5>
        <hr />
        <h5 className="card-text text-center">
          <label htmlFor="password" className="mb-1">
            Password
          </label>
          <input
            className="form-control mb-3"
            onChange={inputHandler}
            value={userFormInfo.password}
            type="password"
            id="password"
            name="password"
          />
        </h5>
        <button className="btn btn-primary me-3" onClick={login}>
          Login
        </button>
      </div>
    </div>
  );
};

export default Login;
