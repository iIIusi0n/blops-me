"use server";

import {cookies} from "next/headers"

export async function getAuthorized() {
    const token = cookies().get("token")?.value;
    if (!token) {
        return false;
    }

    const resp = await fetch(`${process.env.APP_API_URL}/auth/verify`, {
        credentials: 'include',
        headers: {
            Cookie: `token=${token}`
        }
    });
    return resp.ok;
}
