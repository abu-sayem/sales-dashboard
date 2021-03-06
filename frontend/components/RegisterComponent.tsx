import React, { Component, SyntheticEvent } from "react";
import { withRouter } from "next/router";
import { WithRouterProps } from "next/dist/client/with-router";
import axios from "axios";

import "./Login.module.css";
import { route } from "next/dist/server/router";
import { redirect } from "next/dist/server/api-utils";

class Register extends Component<WithRouterProps> {
  first_name = "zZ";
  last_name = ";";
  password = "";
  email = "";
  password_confirm = "";

  state = {
    redirect: false,
  };

  submit = async (e: SyntheticEvent) => {
    e.preventDefault();
    await axios.post("register", {
      first_name: this.first_name,
      last_name: this.last_name,
      password: this.password,
      email: this.email,
      password_confirm: this.password_confirm,
    });

    this.setState({
      redirect: true,
    });
  };

  render() {
    const { router } = this.props;

    if (this.state.redirect) {
      router.push("/");
    }

    return (
      <main className="form-signin d-flex align-items-center justify-content-center my-auto">
        <form onSubmit={this.submit}>
          <h1 className="h3 mb-3 fw-normal">Please Register</h1>
          <input
            className="form-control"
            id="floatingInput"
            placeholder="First Name"
            onChange={(e) => (this.first_name = e.target.value)}
          />
          <input
            className="form-control"
            id="floatingInput"
            placeholder="Last Name"
            onChange={(e) => (this.last_name = e.target.value)}
          />
          <input
            className="form-control"
            id="floatingInput"
            placeholder="name@example.com"
            onChange={(e) => (this.email = e.target.value)}
          />

          <input
            type="password"
            className="form-control"
            id="floatingPassword"
            placeholder="Password"
            onChange={(e) => (this.password = e.target.value)}
          />

          <input
            type="password"
            className="form-control"
            id="floatingPassword"
            placeholder="Confirm Password"
            onChange={(e) => (this.password_confirm = e.target.value)}
          />

          <div className="checkbox mb-3">
            <label>
              <input type="checkbox" value="remember-me" /> Remember me
            </label>
          </div>
          <button className="w-100 btn btn-lg btn-primary" type="submit">
            Sign in
          </button>
          <p className="mt-5 mb-3 text-muted">&copy; 2017???2021</p>
        </form>
      </main>
    );
  }
}

export default withRouter(Register);
