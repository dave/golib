package dwarf

type PackageSession struct {
	abbrevs  [38]dwAbbrev
	logDwarf bool

	sevenbits [128]byte
}

func NewPackageSession() *PackageSession {
	psess := &PackageSession{}
	psess.sevenbits = [...]byte{
		0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
		0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
		0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
		0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
		0x40, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4a, 0x4b, 0x4c, 0x4d, 0x4e, 0x4f,
		0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x5a, 0x5b, 0x5c, 0x5d, 0x5e, 0x5f,
		0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
		0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f,
	}
	psess.abbrevs = [DW_NABRV]dwAbbrev{

		{0, 0, []dwAttrForm{}},

		{
			DW_TAG_compile_unit,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_language, DW_FORM_data1},
				{DW_AT_stmt_list, DW_FORM_sec_offset},
				{DW_AT_low_pc, DW_FORM_addr},
				{DW_AT_ranges, DW_FORM_sec_offset},
				{DW_AT_comp_dir, DW_FORM_string},
				{DW_AT_producer, DW_FORM_string},
			},
		},

		{
			DW_TAG_subprogram,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_low_pc, DW_FORM_addr},
				{DW_AT_high_pc, DW_FORM_addr},
				{DW_AT_frame_base, DW_FORM_block1},
				{DW_AT_decl_file, DW_FORM_data4},
				{DW_AT_external, DW_FORM_flag},
			},
		},

		{
			DW_TAG_subprogram,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_inline, DW_FORM_data1},
				{DW_AT_external, DW_FORM_flag},
			},
		},

		{
			DW_TAG_subprogram,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_low_pc, DW_FORM_addr},
				{DW_AT_high_pc, DW_FORM_addr},
				{DW_AT_frame_base, DW_FORM_block1},
			},
		},

		{
			DW_TAG_inlined_subroutine,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_low_pc, DW_FORM_addr},
				{DW_AT_high_pc, DW_FORM_addr},
				{DW_AT_call_file, DW_FORM_data4},
				{DW_AT_call_line, DW_FORM_udata},
			},
		},

		{
			DW_TAG_inlined_subroutine,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_ranges, DW_FORM_sec_offset},
				{DW_AT_call_file, DW_FORM_data4},
				{DW_AT_call_line, DW_FORM_udata},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_location, DW_FORM_block1},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_external, DW_FORM_flag},
			},
		},

		{
			DW_TAG_constant,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_const_value, DW_FORM_sdata},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_decl_line, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_block1},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_decl_line, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_sec_offset},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_decl_line, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_block1},
			},
		},

		{
			DW_TAG_variable,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_sec_offset},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_variable_parameter, DW_FORM_flag},
				{DW_AT_decl_line, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_block1},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_variable_parameter, DW_FORM_flag},
				{DW_AT_decl_line, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_sec_offset},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_variable_parameter, DW_FORM_flag},
				{DW_AT_type, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_block1},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_abstract_origin, DW_FORM_ref_addr},
				{DW_AT_location, DW_FORM_sec_offset},
			},
		},

		{
			DW_TAG_lexical_block,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_ranges, DW_FORM_sec_offset},
			},
		},

		{
			DW_TAG_lexical_block,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_low_pc, DW_FORM_addr},
				{DW_AT_high_pc, DW_FORM_addr},
			},
		},

		{
			DW_TAG_member,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_data_member_location, DW_FORM_udata},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_go_embedded_field, DW_FORM_flag},
			},
		},

		{
			DW_TAG_formal_parameter,
			DW_CHILDREN_no,

			[]dwAttrForm{
				{DW_AT_type, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_unspecified_parameters,
			DW_CHILDREN_no,
			[]dwAttrForm{},
		},

		{
			DW_TAG_subrange_type,
			DW_CHILDREN_no,

			[]dwAttrForm{
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_count, DW_FORM_udata},
			},
		},

		{
			DW_TAG_unspecified_type,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
			},
		},

		{
			DW_TAG_base_type,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_encoding, DW_FORM_data1},
				{DW_AT_byte_size, DW_FORM_data1},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_array_type,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_byte_size, DW_FORM_udata},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_typedef,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
				{DW_AT_go_elem, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_subroutine_type,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_byte_size, DW_FORM_udata},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_typedef,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_typedef,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
				{DW_AT_go_key, DW_FORM_ref_addr},
				{DW_AT_go_elem, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_pointer_type,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_pointer_type,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
			},
		},

		{
			DW_TAG_structure_type,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_byte_size, DW_FORM_udata},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
				{DW_AT_go_elem, DW_FORM_ref_addr},
			},
		},

		{
			DW_TAG_structure_type,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_byte_size, DW_FORM_udata},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_structure_type,
			DW_CHILDREN_yes,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_byte_size, DW_FORM_udata},
				{DW_AT_go_kind, DW_FORM_data1},
				{DW_AT_go_runtime_type, DW_FORM_addr},
			},
		},

		{
			DW_TAG_typedef,
			DW_CHILDREN_no,
			[]dwAttrForm{
				{DW_AT_name, DW_FORM_string},
				{DW_AT_type, DW_FORM_ref_addr},
			},
		},
	}
	return psess
}
