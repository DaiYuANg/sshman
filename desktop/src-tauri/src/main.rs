// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

use sqlx::migrate::Migrator;
use sqlx::SqlitePool;
static MIGRATOR: Migrator = sqlx::migrate!();

#[tokio::main]
async fn main() {
    let pool = SqlitePool::connect("sqlite::memory:").await.expect("database connection");
    MIGRATOR.run(&pool).await.expect("run");
    desktop_lib::run()
}
