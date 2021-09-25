import axios from "axios";
import { useEffect, useState } from "react";
//import {Redirect} from 'react-router-dom'

import { useRouter } from "next/router";

import Menu from "./Menu";
import Nav from "./Nav";

const Layout = ({ children }: { children: any }) => {
  
  const [redirect, setRedirect] = useState(false)
  const router = useRouter()

  useEffect(() => {
    (async () => {
      try {
      const {data} = await axios.get("user");
      }
      catch(e){
        setRedirect(true)
      }
    })();
  }, []);

  if (redirect) {
     router.push("/login");
  }

  return (
    <>
      <div className="container-fluid">
        <Nav />

        <div className="container-fluid">
          <div className="row">
            <Menu />
            <main className="col-md-9 ms-sm-auto col-lg-10 px-md-4">
              {children}
            </main>
          </div>
        </div>
      </div>
    </>
  );
};

export default Layout;
