import axios from "axios";
import React, { SyntheticEvent, useEffect, useState } from "react";
import { useRouter } from "next/router";
import Layout from "../../../components/Layout";
import { Role } from "../../../models/role";

const userCreate = () => {
  const router = useRouter();

  const [first_name, setfirstName] = useState('');
  const [last_name, setlastName] = useState('')
  const [email, setemail] = useState('');
  const [role_id, setRoleId] = useState(0);
  const [roles, setRoles] = useState([]);
  const [redirect, setredirect] = useState(false);

  useEffect(() => {
    (
      async () => {
        const {data} = await axios.get('roles');
        setRoles(data);
        console.log(data)
      }
    )()
  },[])

  const formSubmit= async (e:SyntheticEvent) => {
    e.preventDefault();
    await axios.post('users', {
      first_name,
      last_name,
      email,
      role_id
    });
    setredirect(true)
    console.log(role_id);
  }

  if (redirect) {
    router.push("/users");
  }

  return (
    <Layout>
      <form onSubmit={formSubmit}>
        <div className="mb-3">
          <label>First Name</label>
          <input
            className="form-control"
            onChange={(e) => setfirstName(e.target.value)}
          />
        </div>
        <div className="mb-3">
          <label>Last Name</label>
          <input
            className="form-control"
            onChange={(e) => setlastName(e.target.value)}
          />
        </div>
        <div className="mb-3">
          <label>Email</label>
          <input
            className="form-control"
            onChange={(e) => setemail(e.target.value)}
          />
        </div>
        <div className="mb-3">
          <label>Role</label>
          <select
            className="form-control"
            onChange={(e) => setRoleId(parseInt(e.target.value))}
          >
            {roles.map((r: Role) => {
              return (
                <option key={r.id} value={r.id}>
                  {r.name}
                </option>
              );
            })}
          </select>
        </div>
        <button className="btn btn-outline-secondary">Save</button>
      </form>
    </Layout>
  );
};

export default userCreate;
