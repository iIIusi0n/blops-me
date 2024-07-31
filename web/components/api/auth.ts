"use server";

import {cookies} from "next/headers";

export async function getAuthorized() {
    const token = cookies().get("token")?.value;
    if (!token) {
        return false;
    }

    const resp = await fetch("http://127.0.0.1:8010/auth/verify", {
        credentials: 'include',
        headers: {
            Cookie: `token=${token}`
        }
    });
    return resp.ok;
}
