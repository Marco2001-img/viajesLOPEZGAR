import React from 'react'
import { Link } from 'react-router-dom';

import userImg from '../../../assets/images/profile/user-1.jpg';

export default function Navbar() {
  return (
    <div>
       <header className="app-header">
       <nav className="navbar navbar-expand-lg navbar-light bg-white fixed-top ">
          <ul className="navbar-nav  ms-auto  d-flex align-items-center me-5">
            <li className="nav-item d-block d-xl-none">
              <a className="nav-link sidebartoggler" id="headerCollapse" href="#">
                <i className="ti ti-menu-2"></i>
              </a>
            </li>
            <li className="nav-item dropdown">
              <a className="nav-link" href="#" id="drop1" data-bs-toggle="dropdown" aria-expanded="false">
                <i className="ti ti-bell"></i>
                <div className="notification bg-primary rounded-circle"></div>
              </a>
              <div className="dropdown-menu dropdown-menu-animate-up" aria-labelledby="drop1">
                <div className="message-body">
                  <a href="#" className="dropdown-item">Item 1</a>
                  <a href="#" className="dropdown-item">Item 2</a>
                </div>
              </div>
            </li>
          </ul>

       <div className="collapse navbar-collapse justify-content-end" id="navbarNav">
  <ul className="navbar-nav flex-row ms-auto align-items-center">
    <li className="nav-item dropdown">
      <a
        className="nav-link"
        href="#"
        id="drop2"
        role="button"
        data-bs-toggle="dropdown"
        aria-expanded="false"
      >
        <img
          src={userImg}
          alt="User"
          width="35"
          height="35"
          className="rounded-circle"
        />
      </a>

      <ul
        className="dropdown-menu dropdown-menu-end dropdown-menu-animate-up"
        aria-labelledby="drop2"
      >
        <li>
          <a className="dropdown-item d-flex align-items-center gap-2" href="#">
            <i className="ti ti-user fs-6"></i>
            <span>My Profile</span>
          </a>
        </li>
        <li>
          <a className="dropdown-item d-flex align-items-center gap-2" href="#">
            <i className="ti ti-mail fs-6"></i>
            <span>My Account</span>
          </a>
        </li>
        <li>
          <a className="dropdown-item d-flex align-items-center gap-2" href="#">
            <i className="ti ti-list-check fs-6"></i>
            <span>My Task</span>
          </a>
        </li>
        <li>
          <Link
            to="/login"
            className="btn btn-outline-primary mx-3 mt-2 d-block"
          >
            Logout
          </Link>
        </li>
      </ul>
    </li>
  </ul>
</div>

        </nav>
      </header>
    </div>
  )
}
