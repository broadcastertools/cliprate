import React from "react";
import { Route, Routes } from "react-router-dom";
import { LoginScreen, VerifyCode } from "./Login";

export const AuthApp = () => {
    return <Routes>
        <Route path="/" element={<LoginScreen />} />
        <Route path="/oauth/callback" element={<VerifyCode />} />
    </Routes>
};
