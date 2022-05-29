import React from "react";

export default function Code({ writable, children }) {
    return (
        <pre className="code">
            <code>{children}</code>
        </pre>
    );
}