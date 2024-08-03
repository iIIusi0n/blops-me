export function encodeString(input: string): string {
    const buffer = Buffer.from(input);

    let base64 = buffer.toString('base64');

    return base64
        .replace(/\+/g, '-')
        .replace(/\//g, '_')
        .replace(/=+$/, '');
}

export function decodeString(input: string): string {
    input = input.replace(/-/g, '+').replace(/_/g, '/');
    const pad = input.length % 4;
    if (pad) {
        if (pad === 1) {
            throw new Error('InvalidLengthError: Input base64url string is the wrong length to determine padding');
        }
        input += new Array(5-pad).join('=');
    }

    const buffer = Buffer.from(input, 'base64');
    return buffer.toString('utf-8');
}
