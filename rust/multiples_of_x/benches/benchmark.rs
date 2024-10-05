extern crate criterion;
use criterion::{black_box, criterion_group, criterion_main, Criterion};
use multiples_of_x::count_by;

fn benchmark_multiples(c: &mut Criterion) {
    c.bench_function("count_by 5x5", |b| b.iter(|| count_by(black_box(5), black_box(5))));
    c.bench_function("count_by 10x10", |b| b.iter(|| count_by(black_box(10), black_box(10))));
    c.bench_function("count_by 20x20", |b| b.iter(|| count_by(black_box(20), black_box(20))));
}

criterion_group!(benches, benchmark_multiples);
criterion_main!(benches);
