package prefetch

func (self *FileInformationWin10) Filenames() []string {
	result := []string{}

	start_of_strings_table := self.FilenameOffset()

	file_metrics_offset := int64(self.FileMetricsOffset())
	for i := uint32(0); i < self.NumberOfFileMetrics(); i++ {
		metric := self.Profile.FileMetricsEntryV30(
			self.Reader, file_metrics_offset)

		filename := ParseUTF16String(
			self.Reader,
			int64(start_of_strings_table+metric.FilenameOffset()),
			int64(metric.FilenameLength()*2))
		result = append(result, filename)
		file_metrics_offset += int64(metric.Size())
	}

	return result
}

func (self *FileInformationVista) Filenames() []string {
	result := []string{}

	start_of_strings_table := self.FilenameOffset()

	file_metrics_offset := int64(self.FileMetricsOffset())
	for i := uint32(0); i < self.NumberOfFileMetrics(); i++ {
		metric := self.Profile.FileMetricsEntryV30(
			self.Reader, file_metrics_offset)

		filename := ParseUTF16String(
			self.Reader,
			int64(start_of_strings_table+metric.FilenameOffset()),
			int64(metric.FilenameLength()*2))
		result = append(result, filename)
		file_metrics_offset += int64(metric.Size())
	}

	return result
}

func (self *FileInformationXP) Filenames() []string {
	result := []string{}

	start_of_strings_table := self.FilenameOffset()

	file_metrics_offset := int64(self.FileMetricsOffset())
	for i := uint32(0); i < self.NumberOfFileMetrics(); i++ {
		metric := self.Profile.FileMetricsEntryV17(
			self.Reader, file_metrics_offset)

		filename := ParseUTF16String(
			self.Reader,
			int64(start_of_strings_table+metric.FilenameOffset()),
			int64(metric.FilenameLength()*2))
		result = append(result, filename)
		file_metrics_offset += int64(metric.Size())
	}

	return result
}
