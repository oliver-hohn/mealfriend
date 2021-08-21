provider "google" {
  project = "mealfriend-323117"
  region  = "europe-west1"
  zone    = "europe-west1-b"
}

resource "google_storage_bucket" "datastore" {
    name = "mealfriend-datastore"
    location = "US-EAST1"
    uniform_bucket_level_access = true
    versioning {
      enabled = true
    }

    force_destroy = true
}

resource "google_service_account" "datastore_account" {
  account_id   = "datastore"
  display_name = "Datastore Account"
}

resource "google_storage_bucket_iam_member" "read_and_write_to_datastore_iam" {
  bucket = google_storage_bucket.datastore.name
  role = "roles/storage.objectAdmin"
  member = "serviceAccount:${google_service_account.datastore_account.email}"

  condition {
    title       = "datastore_only"
    expression  = <<CONDITION
      resource.type == "storage.googleapis.com/Bucket" && resource.name == "${google_storage_bucket.datastore.url}"
    CONDITION
  }
}

resource "google_service_account_key" "datastore_account_key" {
  service_account_id = google_service_account.datastore_account.name
}
