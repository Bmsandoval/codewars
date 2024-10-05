// extern crate criterion;
// use criterion::{black_box, criterion_group, criterion_main, Criterion};
// use are_you_playing_banjo::are_you_playing_banjo;
//
// fn benchmark_banjo(c: &mut Criterion) {
//     c.bench_function("are_you_playing_banjo 5x5", |b| b.iter(|| are_you_playing_banjo(black_box(5))));
//     c.bench_function("are_you_playing_banjo 10x10", |b| b.iter(|| are_you_playing_banjo(black_box(10))));
//     c.bench_function("are_you_playing_banjo 20x20", |b| b.iter(|| are_you_playing_banjo(black_box(20))));
// }
//
// criterion_group!(benches, benchmark_banjo);
// criterion_main!(benches);
