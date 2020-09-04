export function isToday(dateString) {
  const timestamp = Date.parse(dateString);

  if (isNaN(timestamp)) {
    return false;
  }

  const date = new Date(dateString);
  const today = new Date();
  return date.getDate() == today.getDate() &&
    date.getMonth() == today.getMonth() &&
    date.getFullYear() == today.getFullYear();
}
