package salts

var AllSalts = func() []string {
    combined := make([]string, 0, len(SaltsLength8)+len(SaltsLength12)+len(SaltsLength16))
    combined = append(combined, SaltsLength8...)
    combined = append(combined, SaltsLength12...)
    combined = append(combined, SaltsLength16...)
    return combined
}()
