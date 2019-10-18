package mod_5

test_allow_5 {
    allow with input as { "candidate": 5 }
}

test_allow_15 {
    allow with input as { "candidate": 15 }
}

test_disallow_with_3 {
    not allow with input as { "candidate": 3 }
}
