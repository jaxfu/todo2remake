import styles from "./Register.module.scss";
import { useState } from "react";
import { type T_APIRESULT_VALID, type T_FORMINFO_REGISTER } from "../../types";
import { apiRequestRegister } from "../../methods/requests";
import { AxiosResponse } from "axios";

interface IProps {
  setUsername: React.Dispatch<React.SetStateAction<string>>
}

const Register: React.FC<IProps> = (props) => {
  const [userFormInfo, setUserFormInfo] = useState<T_FORMINFO_REGISTER>({
    username: "",
    firstPassword: "",
    secondPassword: "",
  });

  function inputHandler(e: React.ChangeEvent<HTMLInputElement>): void {
    setUserFormInfo((userFormInfo: T_FORMINFO_REGISTER) => {
      return {
        ...userFormInfo,
        [e.target.name]: e.target.value,
      };
    });
  }

  async function register(): Promise<void> {
    const res: AxiosResponse<T_APIRESULT_VALID> = await apiRequestRegister(
      userFormInfo
    );
    console.log(`register valid: ${res.data.valid}`);
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
            value={userFormInfo.firstPassword}
            type="password"
            id="password"
            name="firstPassword"
          />
        </h5>
        <h5 className="card-text text-center">
          <label htmlFor="secondPassword" className="mb-1">
            Confirm Password
          </label>
          <input
            className="form-control mb-3"
            onChange={inputHandler}
            value={userFormInfo.secondPassword}
            type="password"
            id="secondPassword"
            name="secondPassword"
          />
        </h5>
        <button className="btn btn-info me-3" onClick={register}>
          Register
        </button>
      </div>
    </div>
  );
};

export default Register;
