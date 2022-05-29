import React from "react";

export default function Select({ callback, active }) {
    const change = (e) => callback(e.target.value);
    const options = [
        "json",
        "yaml",
        "csv",
        "xml",
    ];

    return (
        <select name="language" onChange={change}>
            {
                options.map(
                    (opt, i) => (
                        opt !== active
                            ? <option value={opt} key={i}>{opt.toUpperCase()}</option>
                            : null
                    )
                )
            }
        </select>
    );
};