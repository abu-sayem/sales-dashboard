import Link from "next/link";
import { useRouter } from "next/router";

const Menu = () => {
  const router = useRouter();
  return (
    <nav
      id="sidebarMenu"
      className="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse"
    >
      <div className="position-sticky pt-3">
        <ul className="nav flex-column">
          <li className="nav-item">
            <Link href="/dashboard">
              <a
                className={`nav-link ${
                  router.pathname == "/dashboard" ? "active" : ""
                }`}
              >
                Dashboard
              </a>
            </Link>
            <Link href="/users">
              <a
                className={`nav-link ${
                  router.pathname == "/users" ? "active" : ""
                }`}
              >
                Users
              </a>
            </Link>
          </li>
        </ul>
      </div>
    </nav>
  );
};

export default Menu;
