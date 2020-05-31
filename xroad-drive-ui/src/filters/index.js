// Format date time. Result YYYY-MM-DD
export function formatDate (dateString) {
  const timestamp = Date.parse(dateString);

  if (isNaN(timestamp)) {
    return '-';
  }

  const date = new Date(dateString);

  return date.toDateString().split(' ').slice(1).join(' ');
}

// Format date time. Result HH:MM.
export function formatHoursMins (dateString) {
  const timestamp = Date.parse(dateString);

  if (isNaN(timestamp)) {
    return '-';
  }

  const date = new Date(dateString);

  return date.getHours().toString().padStart(2, '0') + ':'
    + date.getMinutes().toString().padStart(2, '0');
}

// Format bytes.
export function formatBytes (bytes) {
  if (typeof bytes !== 'number' || isNaN(bytes)) {
    throw new TypeError('Expected a number');
  }

  const UNITS = ['B', 'KB', 'MB', 'GB', 'TB'];

  if (bytes < 0) {
    return 0 + ' ' + UNITS[0];
  } else if (bytes < 1) {
    return bytes + ' ' + UNITS[0];
  }

  const exponent = Math.min(Math.floor(Math.log(bytes) / Math.log(1000)), UNITS.length - 1);
  const number = (bytes / Math.pow(1000, exponent)).toFixed(0);

  return number + ' ' + UNITS[exponent];
}