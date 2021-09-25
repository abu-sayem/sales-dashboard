import React, { Component, useState, SyntheticEvent } from "react";

import { useRouter } from "next/router";

import axios from "axios";

const LoginComponent = () => {
  const router = useRouter();

  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [redirect, setRedirect] = useState(false);

  const submit = async (e) => {
    e.preventDefault();

    const { data } = await axios.post(
      "login",
      {
        email,
        password,
      }
    );

    setRedirect(true);
  };

  if (redirect) {
    router.push("/");
  }

  return (
    <main className="form-signin d-flex align-items-center justify-content-center my-auto">
      <form onSubmit={submit}>
        <h1 className="h3 mb-3 fw-normal">Please Sig n In</h1>
        <input
          className="form-control"
          id="floatingInput"
          placeholder="name@example.com"
          onChange={(e) => setEmail(e.target.value)}
        />

        <input
          type="password"
          className="form-control"
          id="floatingPassword"
          placeholder="Password"
          onChange={(e) => setPassword(e.target.value)}
        />

        <div className="checkbox mb-3">
          <label>
            <input type="checkbox" value="remember-me" /> Remember me
          </label>
        </div>
        <button className="w-100 btn btn-lg btn-primary" type="submit">
          Sign In
        </button>
        <p className="mt-5 mb-3 text-muted">&copy; 2017–2021</p>
      </form>
    </main>
  );
};

export default LoginComponent;
