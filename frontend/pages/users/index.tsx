import React, { useEffect, useState } from "react";
import Link from "next/link";

import Nav from "../../components/Nav";
import Menu from "../../components/Menu";
import Layout from "../../components/Layout";
import axios from "axios";
import Paginator from "../../components/paginator";

import { User } from "../../models/user";

const Users = () => {
  const [users, setUsers] = useState([]);
  const [page, setPage] = useState(0);
  const [lastPage, setLastPage] = useState(0);

  useEffect(() => {
    (async () => {
      const { data } = await axios.get(`users?page=${page}`);
      setUsers(data.data);
      setLastPage(data.meta.last_page);
    })();
  }, [page]);

  const handleDelete = async (id: number) => {
    if (window.confirm("Are you want to delete this message?")) {
      // await axios.delete(`users/${id}`);

      setUsers(
        users.filter((u: User) => {
          u.id !== id;
        })
      );
    }
  };

  return (
    <Layout>
      <Link href="/users/create">
      <a href="" className="btn btn-sm btn-outline-secondary">
        Add
      </a>
      </Link>
      <div className="table-responsive">
        <table className="table table-striped table-sm">
          <thead>
            <tr>
              <th scope="col">#</th>
              <th scope="col">Name</th>
              <th scope="col">Email</th>
              <th scope="col">Role</th>
              <th scope="col">Action</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user: User) => {
              return (
                <tr key={user.id}>
                  <td>{user.id}</td>
                  <td>
                    {user.first_name} {user.last_name}
                  </td>
                  <td>{user.email}</td>
                  <td>{user.role.name}</td>
                  <td>
                    <div className="btn-group mr-2">
                      <a
                        onClick={() => handleDelete(user.id)}
                        href=""
                        className="btn btn-sm btn-outline-secondary"
                      >
                        Delete
                      </a>
                    </div>
                  </td>
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>

      <Paginator page={page} lastPage={lastPage} pageChange={setPage} />
    </Layout>
  );
};

export default Users;
