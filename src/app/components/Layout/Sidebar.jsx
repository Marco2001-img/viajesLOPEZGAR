import React from 'react'
import logo from '../../../assets/images/logos/logo_sinFondo.png';
import { Link } from 'react-router-dom';


export default function Sidebar() {
  return (
    <div>
       <aside className="left-sidebar">
        <div>
          <div className="brand-logo d-flex align-items-center justify-content-between">
            <Link to="/" className="text-nowrap logo-img">
              <img src={logo} alt="Logo" className='w-75'/>
            </Link>
            <div className="close-btn d-xl-none d-block sidebartoggler cursor-pointer" id="sidebarCollapse">
              <i className="ti ti-x fs-6"></i>
            </div>
          </div>

          {/* Sidebar Navigation */}
          <nav className="sidebar-nav scroll-sidebar" data-simplebar="">
            <ul id="sidebarnav">
              <li className="nav-small-cap">
                <iconify-icon icon="solar:menu-dots-linear" className="nav-small-cap-icon fs-4"></iconify-icon>
                <span className="hide-menu">Home</span>
              </li>

              <li className="sidebar-item">
                <Link className="sidebar-link" to="/" aria-expanded="false">
                  <i className="ti ti-atom"></i>
                  <span className="hide-menu">Dashboard</span>
                </Link>
              </li>

              <li className="sidebar-item">
                <Link className="sidebar-link justify-content-between" to="#" aria-expanded="false">
                  <div className="d-flex align-items-center gap-3">
                    <span className="d-flex">
                      <i className="ti ti-aperture"></i>
                    </span>
                    <span className="hide-menu">Analytical</span>
                  </div>
                </Link>
              </li>

              <li className="sidebar-item">
                <Link className="sidebar-link justify-content-between" to="#" aria-expanded="false">
                  <div className="d-flex align-items-center gap-3">
                    <span className="d-flex">
                      <i className="ti ti-shopping-cart"></i>
                    </span>
                    <span className="hide-menu">eCommerce</span>
                  </div>
                </Link>
              </li>

              <li className="sidebar-item">
                <Link className="sidebar-link justify-content-between has-arrow" to="#" aria-expanded="false">
                  <div className="d-flex align-items-center gap-3">
                    <span className="d-flex">
                      <i className="ti ti-layout-grid"></i>
                    </span>
                    <span className="hide-menu">Front Pages</span>
                  </div>
                </Link>

                <ul aria-expanded="false" className="collapse first-level">
                  <li className="sidebar-item">
                    <Link className="sidebar-link justify-content-between" to="#">
                      <div className="d-flex align-items-center gap-3">
                        <div className="round-16 d-flex align-items-center justify-content-center">
                          <i className="ti ti-circle"></i>
                        </div>
                        <span className="hide-menu">Homepage</span>
                      </div>
                    </Link>
                  </li>
                  <li className="sidebar-item">
                    <Link className="sidebar-link justify-content-between" to="#">
                      <div className="d-flex align-items-center gap-3">
                        <div className="round-16 d-flex align-items-center justify-content-center">
                          <i className="ti ti-circle"></i>
                        </div>
                        <span className="hide-menu">About Us</span>
                      </div>
                    </Link>
                  </li>
                  {/* Agrega los demás items del menú de la misma forma */}
                </ul>
              </li>
            </ul>
          </nav>
        </div>
      </aside>
    </div>
  )
}
