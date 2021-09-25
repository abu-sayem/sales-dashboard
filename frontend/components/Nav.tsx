import React, { useEffect, useState } from "react";

import Link from "next/link";

import axios from "axios";
import { User } from "../models/user";

const Nav = () => {
  const [user, setUser] = useState(new User());

  useEffect(() => {
    (async () => {
      const { data } = await axios.get("user");
      setUser(new User(
        data.id,
        data.first_name,
        data.last_name,
        data.email,
      ));
    })();
  }, []);

  const logout = async () => {
    await axios.post("logout", {});
  };

  return (
    <header className="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
      <a className="navbar-brand col-md-3 col-lg-2 me-0 px-3" href="#">
        Company name
      </a>
      <button
        className="navbar-toggler position-absolute d-md-none collapsed"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#sidebarMenu"
        aria-controls="sidebarMenu"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon"></span>
      </button>
      <input
        className="form-control form-control-dark w-100"
        type="text"
        placeholder="Search"
        aria-label="Search"
      />
      <div className="navbar-nav ms-auto ">
        <div className="nav-item text-nowrap d-flex">
          <Link href="/dashboard">
            <a className="nav-link px-3" href="#">
              {user.name}
            </a>
          </Link>

          <div className="nav-item text-nowrap">
            <Link href="/login">
              <a onClick={logout} className="nav-link px-3" href="#">
                Sign out
              </a>
            </Link>
          </div>
        </div>
      </div>
    </header>
  );
};

export default Nav;
