"use server";

import {cookies} from "next/headers";

export async function getAuthorized() {
    const token = cookies().get("token")?.value;
    if (!token) {
        return false;
    }

    const resp = await fetch("http://localhost:8010/auth/verify", {
        credentials: 'include',
        headers: {
            Cookie: `token=${token}`
        }
    });
    return resp.ok;
}
