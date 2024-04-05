package config

const API_DNS = "localhost"
const API_PORT = ":8000"
const API_PROTOCOL = "http://"
const API_UPDATE_URL = API_URL + "/upload"
const API_LIST_URL = API_URL + "/list"
const API_GET_URL = API_URL + "/request/{id}"
const API_DELETE_URL = API_URL + "/audio/{id}"
const API_URL = API_PROTOCOL + API_DNS + API_PORT
const API_SEARCH_URL = API_URL + "/search?q={text}"
