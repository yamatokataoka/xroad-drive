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

// Get date group index.
export function getDateGroupIndex (dateString) {

  const date = new Date(dateString);

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
    return 0;
  }

  else if (date > yesterday) {
    return 1;
  }

  else if (date > thisWeek) {
    return 2;
  }

  else if (date > lastWeek) {
    return 3;
  }

  else if (date > thisMonth) {
    return 4;
  }

  else if (date > lastMonth) {
    return 5;
  }

  else if (date > thisYear) {
    return 6;
  }

  else {
    return 7;
  }
}

// Get date group.
export function getDateGroupByIndex (index) {

  const DateGroups = ['Today', 'Yesterday', 'Earlier This Week', 'Last Week', 'Earlier This Month', 'Last Month', 'Earlier This Year', 'Older'];
  return DateGroups[index];
}