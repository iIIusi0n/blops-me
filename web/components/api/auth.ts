"use server";

import {cookies} from "next/headers";
import {INTERNAL_API_URL} from "@/components/api/config";

export async function getAuthorized() {
    const token = cookies().get("token")?.value;
    if (!token) {
        return false;
    }

    const resp = await fetch(`${INTERNAL_API_URL}/auth/verify`, {
        credentials: 'include',
        headers: {
            Cookie: `token=${token}`
        }
    });
    return resp.ok;
}
