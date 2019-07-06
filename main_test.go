package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var files = []string{
	"2019/05/18", "20190518_115208.jpg", "/9j/4AAQSkZJRgABAQAASABIAAD/4QJqRXhpZgAATU0AKgAAAAgACQEPAAIAAAAIAAAAegEQAAIAAAAJAAAAggESAAMAAAABAAEAAAEaAAUAAAABAAAAjAEbAAUAAAABAAAAlAEoAAMAAAABAAIAAAExAAIAAAAOAAAAnAEyAAIAAAAUAAAAqodpAAQAAAABAAAAvgAAAABzYW1zdW5nAFNNLUc5NzNGAAAAAABIAAAAAQAAAEgAAAABRzk3M0ZYWFUxQVNENQAyMDE5OjA1OjE4IDExOjUyOjA4AAAXgpoABQAAAAEAAAHYgp0ABQAAAAEAAAHgiCIAAwAAAAEAAgAAiCcAAwAAAAEAMgAAkAAABwAAAAQwMjIwkAMAAgAAABQAAAHokAQAAgAAABQAAAH8kgIABQAAAAEAAAIQkgMACgAAAAEAAAIYkgQACgAAAAEAAAIgkgUABQAAAAEAAAIokgcAAwAAAAEAAgAAkgkAAwAAAAEAAAAAkgoABQAAAAEAAAIwoAEAAwAAAAEAAQAAoAIABAAAAAEAAAABoAMABAAAAAEAAAABpAIAAwAAAAEAAAAApAMAAwAAAAEAAAAApAQABQAAAAEAAAI4pAUAAwAAAAEADQAApAYAAwAAAAEAAAAApCAAAgAAACEAAAJAAAAAAAAAAAEAAATgAAAACwAAAAUyMDE5OjA1OjE4IDExOjUyOjA4ADIwMTk6MDU6MTggMTE6NTI6MDgAAAAA4wAAAGQAAASHAAAAMgAAAAAAAAABAAAAPwAAABkAAAAJAAAABQAAAAEAAAABYzcyY2Y1MDBmZmU1MDc1MzAwMDAwMDAwMDAwMDAwMDAAAP/hChVodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6cGhvdG9zaG9wPSJodHRwOi8vbnMuYWRvYmUuY29tL3Bob3Rvc2hvcC8xLjAvIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHBob3Rvc2hvcDpEYXRlQ3JlYXRlZD0iMjAxOS0wNS0xOFQxMTo1MjowOCIgeG1wOk1vZGlmeURhdGU9IjIwMTktMDUtMThUMTE6NTI6MDgiIHhtcDpDcmVhdGVEYXRlPSIyMDE5LTA1LTE4VDExOjUyOjA4IiB4bXA6Q3JlYXRvclRvb2w9Ikc5NzNGWFhVMUFTRDUiLz4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8P3hwYWNrZXQgZW5kPSJ3Ij8+AP/tAHhQaG90b3Nob3AgMy4wADhCSU0EBAAAAAAAPxwBWgADGyVHHAIAAAIAAhwCPwAGMTE1MjA4HAI+AAgyMDE5MDUxOBwCNwAIMjAxOTA1MTgcAjwABjExNTIwOAA4QklNBCUAAAAAABDJNUu9gMuxWWtnhmodZI/z/8AAEQgAAQABAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/bAEMAAwICAwICAwMDAwQDAwQFCAUFBAQFCgcHBggMCgwMCwoLCw0OEhANDhEOCwsQFhARExQVFRUMDxcYFhQYEhQVFP/bAEMBAwQEBQQFCQUFCRQNCw0UFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFP/dAAQAAf/aAAwDAQACEQMRAD8A+KqKKK8I8Q//2Q==",
	"2019/05/25", "20190525_163849.jpg", "/9j/4AAQSkZJRgABAQAASABIAAD/4QI2RXhpZgAATU0AKgAAAAgACQEPAAIAAAAIAAAAegEQAAIAAAAJAAAAggESAAMAAAABAAEAAAEaAAUAAAABAAAAjAEbAAUAAAABAAAAlAEoAAMAAAABAAIAAAExAAIAAAAHAAAAnAEyAAIAAAAUAAAApIdpAAQAAAABAAAAuAAAAABzYW1zdW5nAFNNLUc5NzNGAAAAAABIAAAAAQAAAEgAAAABR29vZ2xlAAAyMDE5OjA1OjI1IDE3OjE5OjAxAAAWgpoABQAAAAEAAAHGgp0ABQAAAAEAAAHOiCIAAwAAAAEAAgAAiCcAAwAAAAEBkAAAkAAABwAAAAQwMjIwkAMAAgAAABQAAAHWkAQAAgAAABQAAAHqkgIABQAAAAEAAAH+kgMACgAAAAEAAAIGkgQACgAAAAEAAAIOkgUABQAAAAEAAAIWkgcAAwAAAAEAAwAAkgkAAwAAAAEAAAAAkgoABQAAAAEAAAIeoAEAAwAAAAEAAQAAoAIABAAAAAEAAAABoAMABAAAAAEAAAABpAIAAwAAAAEAAAAApAMAAwAAAAEAAAAApAQABQAAAAEAAAImpAUAAwAAAAEAGgAApAYAAwAAAAEAAAAAAAAAAAAAAAEAAAAyAAAADAAAAAUyMDE5OjA1OjI1IDE2OjM4OjQ5ADIwMTk6MDU6MjUgMTY6Mzg6NDkAAAAAPwAAABkAAADbAAAAZAAAAAAAAAABAAAAPwAAABkAAABsAAAAGQAAAAEAAAAB/+EKDmh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8APD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNS40LjAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOnBob3Rvc2hvcD0iaHR0cDovL25zLmFkb2JlLmNvbS9waG90b3Nob3AvMS4wLyIgeG1wOkNyZWF0b3JUb29sPSJHb29nbGUiIHhtcDpNb2RpZnlEYXRlPSIyMDE5LTA1LTI1VDE3OjE5OjAxIiB4bXA6Q3JlYXRlRGF0ZT0iMjAxOS0wNS0yNVQxNjozODo0OSIgcGhvdG9zaG9wOkRhdGVDcmVhdGVkPSIyMDE5LTA1LTI1VDE2OjM4OjQ5Ii8+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPD94cGFja2V0IGVuZD0idyI/PgD/7QB4UGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAAD8cAVoAAxslRxwCAAACAAIcAj8ABjE2Mzg0ORwCPgAIMjAxOTA1MjUcAjcACDIwMTkwNTI1HAI8AAYxNjM4NDkAOEJJTQQlAAAAAAAQL4wR6DjajGbUVF87WHpOBv/AABEIAAEAAQMBIgACEQEDEQH/xAAfAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgv/xAC1EAACAQMDAgQDBQUEBAAAAX0BAgMABBEFEiExQQYTUWEHInEUMoGRoQgjQrHBFVLR8CQzYnKCCQoWFxgZGiUmJygpKjQ1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4eLj5OXm5+jp6vHy8/T19vf4+fr/xAAfAQADAQEBAQEBAQEBAAAAAAAAAQIDBAUGBwgJCgv/xAC1EQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/3QAEAAH/2gAMAwEAAhEDEQA/APsKiiivMPTP/9k=",
	"2019/05/26", "20190526_182653.jpg", "/9j/4AAQSkZJRgABAQAASABIAAD/4QJqRXhpZgAATU0AKgAAAAgACQEPAAIAAAAIAAAAegEQAAIAAAAJAAAAggESAAMAAAABAAEAAAEaAAUAAAABAAAAjAEbAAUAAAABAAAAlAEoAAMAAAABAAIAAAExAAIAAAAOAAAAnAEyAAIAAAAUAAAAqodpAAQAAAABAAAAvgAAAABzYW1zdW5nAFNNLUc5NzNGAAAAAABIAAAAAQAAAEgAAAABRzk3M0ZYWFUxQVNFNQAyMDE5OjA1OjI2IDE4OjI2OjUzAAAXgpoABQAAAAEAAAHYgp0ABQAAAAEAAAHgiCIAAwAAAAEAAgAAiCcAAwAAAAEAMgAAkAAABwAAAAQwMjIwkAMAAgAAABQAAAHokAQAAgAAABQAAAH8kgIABQAAAAEAAAIQkgMACgAAAAEAAAIYkgQACgAAAAEAAAIgkgUABQAAAAEAAAIokgcAAwAAAAEAAgAAkgkAAwAAAAEAAAAAkgoABQAAAAEAAAIwoAEAAwAAAAEAAQAAoAIABAAAAAEAAAABoAMABAAAAAEAAAABpAIAAwAAAAEAAAAApAMAAwAAAAEAAAAApAQABQAAAAEAAAI4pAUAAwAAAAEAGgAApAYAAwAAAAEAAAAApCAAAgAAACEAAAJAAAAAAAAAAAEAAACfAAAADAAAAAUyMDE5OjA1OjI2IDE4OjI2OjUzADIwMTk6MDU6MjYgMTg6MjY6NTMAAAAAPwAAABkAAALXAAAAMgAAAAAAAAABAAAAPwAAABkAAABsAAAAGQAAAAEAAAABM2Y5NzUxYWRlZDNiMzk4YTAwMDAwMDAwMDAwMDAwMDAAAP/hChVodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6cGhvdG9zaG9wPSJodHRwOi8vbnMuYWRvYmUuY29tL3Bob3Rvc2hvcC8xLjAvIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHBob3Rvc2hvcDpEYXRlQ3JlYXRlZD0iMjAxOS0wNS0yNlQxODoyNjo1MyIgeG1wOk1vZGlmeURhdGU9IjIwMTktMDUtMjZUMTg6MjY6NTMiIHhtcDpDcmVhdGVEYXRlPSIyMDE5LTA1LTI2VDE4OjI2OjUzIiB4bXA6Q3JlYXRvclRvb2w9Ikc5NzNGWFhVMUFTRTUiLz4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8P3hwYWNrZXQgZW5kPSJ3Ij8+AP/tAHhQaG90b3Nob3AgMy4wADhCSU0EBAAAAAAAPxwBWgADGyVHHAIAAAIAAhwCPwAGMTgyNjUzHAI+AAgyMDE5MDUyNhwCNwAIMjAxOTA1MjYcAjwABjE4MjY1MwA4QklNBCUAAAAAABBL/+hmpwl+MKlIdIOCafPK/8AAEQgAAQABAwEiAAIRAQMRAf/EAB8AAAEFAQEBAQEBAAAAAAAAAAABAgMEBQYHCAkKC//EALUQAAIBAwMCBAMFBQQEAAABfQECAwAEEQUSITFBBhNRYQcicRQygZGhCCNCscEVUtHwJDNicoIJChYXGBkaJSYnKCkqNDU2Nzg5OkNERUZHSElKU1RVVldYWVpjZGVmZ2hpanN0dXZ3eHl6g4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2drh4uPk5ebn6Onq8fLz9PX29/j5+v/EAB8BAAMBAQEBAQEBAQEAAAAAAAABAgMEBQYHCAkKC//EALURAAIBAgQEAwQHBQQEAAECdwABAgMRBAUhMQYSQVEHYXETIjKBCBRCkaGxwQkjM1LwFWJy0QoWJDThJfEXGBkaJicoKSo1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoKDhIWGh4iJipKTlJWWl5iZmqKjpKWmp6ipqrKztLW2t7i5usLDxMXGx8jJytLT1NXW19jZ2uLj5OXm5+jp6vLz9PX29/j5+v/bAEMAAwICAwICAwMDAwQDAwQFCAUFBAQFCgcHBggMCgwMCwoLCw0OEhANDhEOCwsQFhARExQVFRUMDxcYFhQYEhQVFP/bAEMBAwQEBQQFCQUFCRQNCw0UFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFP/dAAQAAf/aAAwDAQACEQMRAD8A+RaKKK/rg/HD/9k=",
	"2019/05/26", "20190526_214716.jpg", "/9j/4AAQSkZJRgABAQAASABIAAD/4QJWRXhpZgAATU0AKgAAAAgACQEPAAIAAAAIAAAAegEQAAIAAAAJAAAAggESAAMAAAABAAEAAAEaAAUAAAABAAAAjAEbAAUAAAABAAAAlAEoAAMAAAABAAIAAAExAAIAAAAOAAAAnAEyAAIAAAAUAAAAqodpAAQAAAABAAAAvgAAAABzYW1zdW5nAFNNLUc5NzNGAAAAAABIAAAAAQAAAEgAAAABRzk3M0ZYWFUxQVNFNQAyMDE5OjA1OjI2IDIxOjQ3OjE2AAAWgpoABQAAAAEAAAHMgp0ABQAAAAEAAAHUiCIAAwAAAAEAAgAAiCcAAwAAAAECgAAAkAAABwAAAAQwMjIwkAMAAgAAABQAAAHckAQAAgAAABQAAAHwkgIABQAAAAEAAAIEkgQACgAAAAEAAAIMkgUABQAAAAEAAAIUkgcAAwAAAAEAAgAAkgkAAwAAAAEAAAAAkgoABQAAAAEAAAIcoAEAAwAAAAEAAQAAoAIABAAAAAEAAAABoAMABAAAAAEAAAABpAIAAwAAAAEAAAAApAMAAwAAAAEAAAAApAQABQAAAAEAAAIkpAUAAwAAAAEAGgAApAYAAwAAAAEAAAAApCAAAgAAACEAAAIsAAAAAAAAAAEAAAAhAAAAAwAAAAIyMDE5OjA1OjI2IDIxOjQ3OjE2ADIwMTk6MDU6MjYgMjE6NDc6MTYAAAAAHQAAABkAAAAAAAAAAQAAAD8AAAAZAAAAbAAAABkAAAABAAAAAWU0OTJmZTJmNGI2Zjk3N2IwMDAwMDAwMDAwMDAwMDAwAAD/4QoVaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLwA8P3hwYWNrZXQgYmVnaW49Iu+7vyIgaWQ9Ilc1TTBNcENlaGlIenJlU3pOVGN6a2M5ZCI/PiA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJYTVAgQ29yZSA1LjQuMCI+IDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+IDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiIHhtbG5zOnBob3Rvc2hvcD0iaHR0cDovL25zLmFkb2JlLmNvbS9waG90b3Nob3AvMS4wLyIgeG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiBwaG90b3Nob3A6RGF0ZUNyZWF0ZWQ9IjIwMTktMDUtMjZUMjE6NDc6MTYiIHhtcDpNb2RpZnlEYXRlPSIyMDE5LTA1LTI2VDIxOjQ3OjE2IiB4bXA6Q3JlYXRlRGF0ZT0iMjAxOS0wNS0yNlQyMTo0NzoxNiIgeG1wOkNyZWF0b3JUb29sPSJHOTczRlhYVTFBU0U1Ii8+IDwvcmRmOlJERj4gPC94OnhtcG1ldGE+ICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPD94cGFja2V0IGVuZD0idyI/PgD/7QB4UGhvdG9zaG9wIDMuMAA4QklNBAQAAAAAAD8cAVoAAxslRxwCAAACAAIcAj8ABjIxNDcxNhwCPgAIMjAxOTA1MjYcAjcACDIwMTkwNTI2HAI8AAYyMTQ3MTYAOEJJTQQlAAAAAAAQSMV4ZzldG9dqChg5B3GvI//AABEIAAEAAQMBIgACEQEDEQH/xAAfAAABBQEBAQEBAQAAAAAAAAAAAQIDBAUGBwgJCgv/xAC1EAACAQMDAgQDBQUEBAAAAX0BAgMABBEFEiExQQYTUWEHInEUMoGRoQgjQrHBFVLR8CQzYnKCCQoWFxgZGiUmJygpKjQ1Njc4OTpDREVGR0hJSlNUVVZXWFlaY2RlZmdoaWpzdHV2d3h5eoOEhYaHiImKkpOUlZaXmJmaoqOkpaanqKmqsrO0tba3uLm6wsPExcbHyMnK0tPU1dbX2Nna4eLj5OXm5+jp6vHy8/T19vf4+fr/xAAfAQADAQEBAQEBAQEBAAAAAAAAAQIDBAUGBwgJCgv/xAC1EQACAQIEBAMEBwUEBAABAncAAQIDEQQFITEGEkFRB2FxEyIygQgUQpGhscEJIzNS8BVictEKFiQ04SXxFxgZGiYnKCkqNTY3ODk6Q0RFRkdISUpTVFVWV1hZWmNkZWZnaGlqc3R1dnd4eXqCg4SFhoeIiYqSk5SVlpeYmZqio6Slpqeoqaqys7S1tre4ubrCw8TFxsfIycrS09TV1tfY2dri4+Tl5ufo6ery8/T19vf4+fr/2wBDAAMCAgMCAgMDAwMEAwMEBQgFBQQEBQoHBwYIDAoMDAsKCwsNDhIQDQ4RDgsLEBYQERMUFRUVDA8XGBYUGBIUFRT/2wBDAQMEBAUEBQkFBQkUDQsNFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBT/3QAEAAH/2gAMAwEAAhEDEQA/AN2iiiv6NPxQ/9k=",
	"2019/05/26", "20190526_104122.mp4", "MAo=",
}

func setUp(t *testing.T) {
	tearDown(t)
	run(t, "mkdir -p testdir")
}

func tearDown(t *testing.T) {
	run(t, "rm -rf testdir")
}

func TestIntegration(t *testing.T) {
	setUp(t)
	for i := 0; i < len(files)-1; i += 3 {
		dir, name, content := files[i], files[i+1], files[i+2]
		err := os.MkdirAll(filepath.Join("testdir", dir), os.ModePerm)
		if err != nil {
			t.Fatal(err)
		}
		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			t.Fatal(err)
		}
		if err = ioutil.WriteFile(filepath.Join("testdir", dir, name), data, 0777); err != nil {
			t.Fatal(err)
		}
	}
	run(t, "go run main.go --dir testdir seal")
	data, err := ioutil.ReadFile("testdir/index.csv")
	if err != nil {
		t.Fatal(err)
	}
	want := `2019/05/18/20190518_115208.jpg,3964,befc025f683d43e6084a11475bef97a2
2019/05/25/20190525_163849.jpg,3905,08b828b7cc3565c7552b62f7bf898acd
2019/05/26/20190526_104122.mp4,2,897316929176464ebc9ad085f31e7284
2019/05/26/20190526_182653.jpg,3965,60f0dc2ed8c723ca267c7f67852e0472
2019/05/26/20190526_214716.jpg,3944,b24c0e40d276a2fff3cd325d933c447a
`
	got := string(data)
	if want != got {
		t.Fatalf("unexpected index.csv: want:\n%s\ngot:\n%s", want, got)
	}
	out := run(t, "go run main.go --dir testdir verify")
	if !strings.Contains(out, "looks sound") {
		t.Fatalf("expected looks sound result, got: %s", out)
	}
	out = run(t, "go run main.go --dir testdir --deep verify")
	if !strings.Contains(out, "is sound") {
		t.Fatalf("expected is sound result, got: %s", out)
	}
	if err = ioutil.WriteFile("testdir/2019/05/26/20190526_104122.mp4", []byte{1, 1}, 0777); err != nil {
		t.Fatal(err)
	}
	out = run(t, "go run main.go --dir testdir verify")
	if !strings.Contains(out, "looks sound") {
		t.Fatalf("expected looks sound result, got: %s", out)
	}
	out = run(t, "go run main.go --dir testdir --deep verify")
	if !strings.Contains(out, "20190526_104122.mp4 is corrupted") {
		t.Fatalf("expected corrupted 20190526_104122.mp4, got: %s", out)
	}
	run(t, "rm testdir/2019/05/26/20190526_104122.mp4")
	out = run(t, "go run main.go --dir testdir verify")
	if !strings.Contains(out, "20190526_104122.mp4 is missing") {
		t.Fatalf("expected 20190526_104122.mp4 missing, got: %s", out)
	}
	run(t, "go run main.go --dir testdir seal")
	out = run(t, "go run main.go --dir testdir --deep verify")
	if !strings.Contains(out, "is sound") {
		t.Fatalf("expected is sound result, got: %s", out)
	}
	for i := 0; i < 5; i += 3 {
		dir, name, content := files[i], files[i+1], files[i+2]
		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			t.Fatal(err)
		}
		if err = ioutil.WriteFile(filepath.Join("testdir", dir, "_"+name), data, 0777); err != nil {
			t.Fatal(err)
		}
	}
	run(t, "go run main.go --dir testdir seal")
	data, err = ioutil.ReadFile("testdir/index.csv")
	if err != nil {
		t.Fatal(err)
	}
	want = `2019/05/18/20190518_115208.jpg,3964,befc025f683d43e6084a11475bef97a2
2019/05/18/_20190518_115208.jpg,3964,befc025f683d43e6084a11475bef97a2
2019/05/25/20190525_163849.jpg,3905,08b828b7cc3565c7552b62f7bf898acd
2019/05/25/_20190525_163849.jpg,3905,08b828b7cc3565c7552b62f7bf898acd
2019/05/26/20190526_182653.jpg,3965,60f0dc2ed8c723ca267c7f67852e0472
2019/05/26/20190526_214716.jpg,3944,b24c0e40d276a2fff3cd325d933c447a
`
	got = string(data)
	if want != got {
		t.Fatalf("unexpected index.csv: want:\n%s\ngot:\n%s", want, got)
	}
	out = run(t, "go run main.go --dir testdir verify")
	if !strings.Contains(out, "looks sound") {
		t.Fatalf("expected looks sound result, got: %s", out)
	}
	out = run(t, "go run main.go --dir testdir --deep verify")
	if !strings.Contains(out, "duplicates") {
		t.Fatalf("expected finding duplicates, got: %s", out)
	}
	tearDown(t)
}

func run(t *testing.T, s string) string {
	fields := strings.Fields(s)
	out, err := exec.Command(fields[0], fields[1:]...).Output()
	if err != nil {
		t.Fatalf("failed on %s: %v", s, err)
	}
	return string(out)
}
