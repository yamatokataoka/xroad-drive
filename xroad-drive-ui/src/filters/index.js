// Format date time. Result MM DD YYYY
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

  const UNITS = ['bytes', 'KB', 'MB', 'GB', 'TB'];

  if (bytes < 0) {
    return 0 + ' ' + UNITS[0];
  } else if (bytes < 1) {
    return bytes + ' ' + UNITS[0];
  }

  const exponent = Math.min(Math.floor(Math.log(bytes) / Math.log(1000)), UNITS.length - 1);
  const number = (bytes / Math.pow(1000, exponent)).toFixed(0);

  return number + ' ' + UNITS[exponent];
}

// Return date group from date string
export function dateGroup (dateString) {

  const date = new Date(dateString);
  const dateGroups = ['Today', 'Yesterday', 'Earlier This Week', 'Last Week', 'Earlier This Month', 'Last Month', 'Earlier This Year', 'Older'];

  const today = new Date();
  today.setHours(0,0,0,0);

  const yesterday = new Date(today);
  yesterday.setDate(today.getDate() - 1);

  const thisWeek = new Date(today);
  const day = today.getDay();
  const diffDate = today.getDate() - day + (day == 0 ? -6:1); // adjust when day is sunday
  thisWeek.setDate(diffDate);

  const lastWeek = new Date(thisWeek);
  lastWeek.setDate(thisWeek.getDate() - 7);

  const thisMonth = new Date(today.getFullYear(), today.getMonth(), 1);

  const lastMonth = new Date(thisMonth);
  lastMonth.setMonth(thisMonth.getMonth() - 1);

  const thisYear = new Date(today.getFullYear(), 0, 1);

  if (date > today) {
    return dateGroups[0];
  }

  else if (date > yesterday) {
    return dateGroups[1];
  }

  else if (date > thisWeek) {
    return dateGroups[2];
  }

  else if (date > lastWeek) {
    return dateGroups[3];
  }

  else if (date > thisMonth) {
    return dateGroups[4];
  }

  else if (date > lastMonth) {
    return dateGroups[5];
  }

  else if (date > thisYear) {
    return dateGroups[6];
  }

  else {
    return dateGroups[7];
  }
}
