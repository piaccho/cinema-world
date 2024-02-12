export default function getTimeFromDate(date: string): string {
    const dateTime = new Date(date);
    const timeString = dateTime.toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' });
    return timeString;
}
