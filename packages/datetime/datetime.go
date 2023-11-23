package datetime

import "time"

type DateTime struct {
	inner time.Time
}

func New(t time.Time) *DateTime {
	return &DateTime{inner: t}
}

func (dt DateTime) StartOfHour() *DateTime {
	return dt.startOfTemplate(dt.inner.Hour(), 0, 0, 0)
}

func (dt DateTime) StartOfDay() *DateTime {
	return dt.startOfTemplate(0, 0, 0, 0)
}

func (dt DateTime) startOfTemplate(hour, min, sec, nsec int) *DateTime {
	date := time.Date(
		dt.inner.Year(),
		dt.inner.Month(),
		dt.inner.Day(),
		hour,
		min,
		sec,
		nsec,
		dt.inner.Location(),
	)

	return New(date)
}

func (dt DateTime) String() string {
	return dt.inner.Format(time.DateTime)
}
