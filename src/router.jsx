import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./app/components/auth/Login";
import Registrar from "./app/components/auth/registrar";
import Inicio from "./app/components/common/Inicio";

export default function AppRouter() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Login />} />
                <Route path="/registrate" element={<Registrar />} />
                <Route path="/dashboard" element={<Inicio />} />
            </Routes>
        </BrowserRouter>
    );
}
