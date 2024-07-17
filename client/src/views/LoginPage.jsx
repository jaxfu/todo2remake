import { useState } from "react";
import { Navigate } from "react-router-dom";

const LoginPage = () => {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [sentInfo, setSentInfo] = useState(false);

  const usernameInputHandler = (e) => {
    setUsername(e.target.value);
  };

  const passwordInputHandler = (e) => {
    setPassword(e.target.value);
  };

  const submitInfo = () => {
    const data = {
      username,
      password,
    };

    try {
      fetch("http://localhost:5000/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      })
        .then((res) => res.json())
        .then(() => {
          setSentInfo(true);
        });
    } catch (e) {
      console.log(`Error: ${e}`);
    }
  };

  return (
    <div className="card mx-auto mt-3 loginCard">
      <div className="card-body">
        <h5 className="card-title">
          <label htmlFor="username" className="mb-1">
            Username
          </label>
          <input
            className="form-control"
            onChange={usernameInputHandler}
            value={username}
            type="text"
            id="username"
          />
        </h5>
        <hr />
        <h5 className="card-text">
          <label htmlFor="password" className="mb-1">
            Password
          </label>
          <input
            className="form-control mb-3"
            onChange={passwordInputHandler}
            value={password}
            type="password"
            id="password"
          />
        </h5>
        <hr />
        <div className="card-text">
          <button className="btn btn-info me-3 text-light" onClick={submitInfo}>
            Login
          </button>
          <button className="btn btn-danger text-light">
            Forgot Password?
          </button>
          {sentInfo && <Navigate replace to="/" />}
        </div>
      </div>
    </div>
  );
};

export default LoginPage;
