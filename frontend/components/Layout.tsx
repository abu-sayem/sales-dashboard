import Menu from "./Menu";
import Nav from "./Nav";

const Layout = ({ children }: {children:any}) => {
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
  
}

export default Layout