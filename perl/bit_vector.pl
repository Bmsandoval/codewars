#!/usr/bin/perl

use strict;

# Function definition
sub sortByBit {
    my $result = 0;
    for (@_) {
        $result |= 1 << $_;
    }
    return $result
}

# ASSERTION TESTS
use Test::More 'no_plan';
Test::More->builder->output('/dev/null');
cmp_ok(sortByBit((1)), 'eq', 2, 'integer equality');
cmp_ok(sortByBit((0,1,2)), 'eq', 7, 'integer equality');
cmp_ok(sortByBit((0,1,2,4)), 'eq', 23, 'integer equality');
cmp_ok(sortByBit((1,2,3,4)), 'eq', 30, 'integer equality');
