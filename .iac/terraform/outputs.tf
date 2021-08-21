output "datastore_account_email" {
    value = google_service_account.datastore_account.email
}

output "datastore_account_key" {
    value = google_service_account_key.datastore_account_key.private_key
    sensitive = true
}
