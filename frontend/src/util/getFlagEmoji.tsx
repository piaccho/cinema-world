export default function getFlagEmoji(countryCode: any) {
    return countryCode.toUpperCase().replace(/./g, (char: any) =>
        String.fromCodePoint(127397 + char.charCodeAt())
    );
}