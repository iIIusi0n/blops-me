import bigInt from "big-integer";

const BASE62_ALPHABET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";
const BASE = bigInt(62);

export function encodeString(input: string): string {
    let bytes = Buffer.from(input, 'utf-8');
    // @ts-ignore
    let num = bigInt.fromArray([...bytes], 256);
    let encoded = '';

    while (num.greater(0)) {
        let { quotient, remainder } = num.divmod(BASE);
        encoded = BASE62_ALPHABET[remainder.toJSNumber()] + encoded;
        num = quotient;
    }

    return encoded;
}

export function decodeString(encoded: string): string {
    let num = bigInt(0);

    for (let char of encoded) {
        num = num.multiply(BASE).add(BASE62_ALPHABET.indexOf(char));
    }

    let bytes = num.toArray(256).value;
    return Buffer.from(bytes).toString('utf-8');
}
