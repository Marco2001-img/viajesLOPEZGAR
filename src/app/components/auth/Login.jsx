import React, { useState } from 'react';
import '../../../assets/css/styles.min.css';
import '../../../assets/libs/jquery/dist/jquery.min.js';
import '../../../assets/libs/bootstrap/dist/js/bootstrap.bundle.min.js';
import { Link } from 'react-router-dom';
import logo from '../../../assets/images/logos/logo_sinFondo.png';

export default function Login() {
  // Estado para inputs
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [remember, setRemember] = useState(true);

  // Manejar envío de formulario
  const handleSubmit = (e) => {
    e.preventDefault(); // Evita recargar la página
    console.log('Email:', email);
    console.log('Password:', password);
    console.log('Remember me:', remember);
    // Aquí puedes agregar lógica de login real con API
  };

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
                    <img src={logo} alt="Logo" className="img-fluid w-50" />
                  </Link>
                  <form onSubmit={handleSubmit}>
                    <div className="mb-3">
                      <label htmlFor="exampleInputEmail1" className="form-label">
                        Username
                      </label>
                      <input
                        type="email"
                        className="form-control"
                        id="exampleInputEmail1"
                        aria-describedby="emailHelp"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                      />
                    </div>
                    <div className="mb-4">
                      <label htmlFor="exampleInputPassword1" className="form-label">
                        Password
                      </label>
                      <input
                        type="password"
                        className="form-control"
                        id="exampleInputPassword1"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                      />
                    </div>
                    <div className="d-flex align-items-center justify-content-between mb-4">
                      <div className="form-check">
                        <input
                          className="form-check-input primary"
                          type="checkbox"
                          id="flexCheckChecked"
                          checked={remember}
                          onChange={() => setRemember(!remember)}
                        />
                        <label className="form-check-label text-dark" htmlFor="flexCheckChecked">
                          Remember this Device
                        </label>
                      </div>
                      <Link className="text-primary fw-bold" to="/">
                        Forgot Password?
                      </Link>
                    </div>
                    <button type="submit" className="btn btn-primary w-100 py-8 fs-4 mb-4 rounded-2">
                      Sign In
                    </button>
                    <div className="d-flex align-items-center justify-content-center">
                      <p className="fs-4 mb-0 fw-bold">New to MaterialM?</p>
                      <Link className="text-primary fw-bold ms-2" to="/registrate">
                        Create an account
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
