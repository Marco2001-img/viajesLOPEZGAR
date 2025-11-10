import React from 'react';
import '../../../assets/css/styles.min.css';
import '../../../assets/libs/jquery/dist/jquery.min.js';
import '../../../assets/libs/bootstrap/dist/js/bootstrap.bundle.min.js';
import { Link } from 'react-router-dom';
import logo from '../../../assets/images/logos/logo_sinFondo.png'; 

export default function Registrar() {
  return (
    <div
      className="page-wrapper"
      id="main-wrapper"
      data-layout="vertical"
      data-navbarbg="skin6"
      data-sidebartype="full"
      data-sidebar-position="fixed"
      data-header-position="fixed"
    >
      <div className="position-relative overflow-hidden text-bg-light min-vh-100 d-flex align-items-center justify-content-center">
        <div className="d-flex align-items-center justify-content-center w-100">
          <div className="row justify-content-center w-100">
            <div className="col-md-8 col-lg-6 col-xxl-3">
              <div className="card mb-0">
                <div className="card-body">
                  <Link to="/" className="text-nowrap logo-img text-center d-block py-3 w-100">
                    <img src={logo} alt="Logo" className="img-fluid w-50"/>
                  </Link>
                  
                  <form>
                    <div className="mb-3">
                      <label htmlFor="exampleInputtext1" className="form-label">Name</label>
                      <input type="text" className="form-control" id="exampleInputtext1" aria-describedby="textHelp" />
                    </div>
                    <div className="mb-3">
                      <label htmlFor="exampleInputEmail1" className="form-label">Email Address</label>
                      <input type="email" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" />
                    </div>
                     <div className="mb-3">
                      <label htmlFor="exampleInputEmail1" className="form-label">telefono</label>
                      <input type="email" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" />
                    </div>
                    <div className="mb-3">
                      <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                      <input type="password" className="form-control" id="exampleInputPassword1" />
                    </div>
                    <Link to="/" className="btn btn-primary w-100 py-2 fs-4 mb-4 rounded-2">
                      Sign Up
                    </Link>
                    <div className="d-flex align-items-center justify-content-center">
                      <p className="fs-4 mb-0 fw-bold">Already have an Account?</p>
                      <Link to="/" className="text-primary fw-bold ms-2">
                        Sign In
                      </Link>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
