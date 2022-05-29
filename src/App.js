import React, { useState } from "react";
import { NavBar, Select } from "./components";

export default function App() {
    const [from, setFrom] = useState("json");
    const [to, setTo] = useState("yaml");

    return (
        <>
            <NavBar />
            <div className="container">
                <div className="from">
                    <Select callback={ setFrom } active={ to } />
                </div>
                <div className="to">
                    <Select callback={ setTo } active={ from } />
                </div>
            </div>
        </>
    );
};